let width = window.innerWidth;
let height = window.innerHeight;
const rectWidth = 250;
const rectHeight = 80;
const space = 60;

let listPerson = [];

const FILL_COLOR = {
    male: '#40739e',
    female: '#8e44ad',
    undefined: '#808e9b',
}

// function to generate a list of "targets" (circles)
function generatePosition(parent, list, widthWrapper, spaceBefore) {
    const leafSize = parent?.leafSize || 1;
    const startPosition = (parent?.x ? parent?.x + (rectWidth / 2) : spaceBefore + (widthWrapper / 2)) - ((leafSize * rectWidth + (leafSize - 1) * space) / 2);
    let totalWrapper = 0;
    return _.flatten(
        _.map(list, item => {
            const childrenLeafSize = item?.leafSize;
            const wrapper = !!parent ? childrenLeafSize * rectWidth + (childrenLeafSize - 1) * space : 0;

            item.parentId = parent?.id || null;
            item.x = startPosition + (wrapper / 2) - (!!parent ? (rectWidth / 2) : 0) + totalWrapper;
            item.y = parent?.y ? (parent?.y + space) + rectHeight : 50;
            totalWrapper += wrapper + space;
            return [item, ...(_.size(item.children) ? generatePosition(item, item.children) : [])];
        })
    );
}

function countTotalChildren(list) {
    return _.map(list, (item, index) => {
        item.children = countTotalChildren(item.children);
        item.leafSize = _.reduce(item.children, (rs, child) => rs + child.leafSize, 0) || 1;
        item.dept = (_.max(_.map(item.children, 'dept')) || 0) + 1;
        return item;
    })
}


function show() {
    listPerson = countTotalChildren(listPerson);
    let listPersonPositions = [];
    let totalWidthWrapper = 0;
    let totalHeightWrapper = 0;
    _.forEach(listPerson, list => {
        const leafSize = list.leafSize;
        const dept = list.dept;
        const widthWrapper = leafSize * rectWidth + (leafSize - 1) * space + 100;
        const heightWrapper = dept * rectHeight + (dept - 1) * space + 100;

        listPersonPositions = _.concat(
            listPersonPositions,
            generatePosition(null, [list], widthWrapper, totalWidthWrapper)
        );

        totalWidthWrapper = totalWidthWrapper + widthWrapper;
        totalHeightWrapper = totalHeightWrapper + heightWrapper;
    })

    width = totalWidthWrapper > window.innerWidth ? totalWidthWrapper : window.innerWidth;
    height = totalHeightWrapper > window.innerHeight ? totalHeightWrapper : window.innerHeight;

    const stage = new Konva.Stage({
        container: 'container',
        width: width,
        height: height,
        draggable: true,
    });

    const scaleBy = 1.5;
    stage.on('wheel', (e) => {
        // stop default scrolling
        e.evt.preventDefault();

        var oldScale = stage.scaleX();
        var pointer = stage.getPointerPosition();

        var mousePointTo = {
            x: (pointer.x - stage.x()) / oldScale,
            y: (pointer.y - stage.y()) / oldScale,
        };

        // how to scale? Zoom in? Or zoom out?
        let direction = e.evt.deltaY > 0 ? 1 : -1;

        // when we zoom on trackpad, e.evt.ctrlKey is true
        // in that case lets revert direction
        if (e.evt.ctrlKey) {
            direction = -direction;
        }

        var newScale = direction > 0 ? oldScale / scaleBy : oldScale * scaleBy;

        stage.scale({ x: newScale, y: newScale });

        var newPos = {
            x: pointer.x - mousePointTo.x * newScale,
            y: pointer.y - mousePointTo.y * newScale,
        };
        stage.position(newPos);
    });

    const layer = new Konva.Layer();
    stage.add(layer);

    listPersonPositions.forEach(child => {
        const group = new Konva.Group({
            // draggable: true,
        });
        const wrapper = new Konva.Rect({
            id: child.id,
            x: child.x,
            y: child.y,
            width: rectWidth,
            height: rectHeight,
            opacity: 0.8,
            fill: FILL_COLOR[child.position_title.gender_id],
            // shadowBlur: 8,
            cornerRadius: 4,
        });
        wrapper.on('pointerdown', function (evt) {
            const allLines = stage.find('Line');
            _.forEach(allLines, line => {
                line.stroke('#707070');
                line.strokeWidth(2);
                line.zIndex(1);
            })
            const shapes = stage.find('Rect');
            _.forEach(shapes, shape => {
                shape.opacity(0.8);
            })

            const currentId = evt.target.attrs.id;
            const nodes = layer.find(`#${currentId}, #${currentId}1`);
            const lines = layer.find(`.line-${currentId}`);
            _.forEach(lines, line => {
                line.strokeWidth(5);
                line.stroke('red');
                line.zIndex(22);
            })
            _.forEach(nodes, node => {
                node.opacity(1);
            })
        })
        const text = new Konva.Text({
            x: child.x,
            y: child.y + 16,
            text: child.birth_name,
            fontSize: 18,
            fontFamily: 'Great Vibes',
            fill: '#fff',
            width: rectWidth,
            align: 'center',
        });
        group.add(wrapper);
        group.add(text);
        layer.add(group);
    });

    listPersonPositions.forEach(child => {
        if (child.parentId) {
            const fromNode = layer.findOne('#' + child.id);
            const toNode = layer.findOne('#' + child.parentId);
            const line = new Konva.Line({
                name: `line-${child.id}`,
                stroke: '#707070',
                strokeWidth: 2,
                lineJoin: 'round',
            });

            const points = getConnectorPoints(
                fromNode.position(),
                toNode.position()
            );
            line.points(points);

            layer.add(line);
        }
    });
}


function getConnectorPoints(from, to) {
    const fromX = from.x + rectWidth / 2;
    const fromY = from.y;

    const toX1 = fromX;
    const toY1 = fromY - (space / 2);

    const toX2 = to.x + rectWidth / 2;
    const toY2 = toY1;

    const toX3 = to.x + rectWidth / 2;
    const toY3 = to.y + rectHeight;

    return [fromX, fromY, toX1, toY1, toX2, toY2, toX3, toY3];
}

fetch('http://52.71.236.216:3000/api/person/all')
    .then((response) => response.json())
    .then((data) => {
        listPerson = buildTree(null, data);
        show();
    });

function buildTree(node, data) {
    if (!node) {
        return _.map(
            _.filter(data, item => !item.father_id),
            item => buildTree(item, data)
        )
    }
    node.children = _.sortBy(
        _.filter(data, item => item.father_id === node.id),
        'no'
    );
    _.forEach(node.children, item => {
        buildTree(item, data);
    })
    node.childrenNestedId = _.flatten(
        _.map(node.children, item => item.id),
        _.flatten(
            _.map(node.children, item => item.childrenNestedId),
        )
    );
    if (!_.size(node.childrenNestedId)) {
        node.childrenNestedId = [node.id];
    }
    return node;
}

