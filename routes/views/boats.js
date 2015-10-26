var keystone = require('keystone');
var async = require('async');

exports = module.exports = function(req, res) {

	var view = new keystone.View(req, res);
	var locals = res.locals;

	// Init locals
	locals.section = 'boats';
	locals.data = {
		boats: []
	};
	view.on('init', function(next) {

		var q = keystone.list('Boat')
			.where('state', 'published');

		q.exec(function(err, results) {
			locals.data.boats = results;
			next(err);
		});

	});

	// Render the view
	view.render('boats');

};
