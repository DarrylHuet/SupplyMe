$('#table1').bootstrapTable({
    url: 'temp.json',
    columns: [{
        field: 'id',
        title: 'ID'
    }, {
        field: 'content',
        title: 'Contents'
    }, {
        field: 'time',
        title: 'Time Stamp'
    }],
})