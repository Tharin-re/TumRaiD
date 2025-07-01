# -*- coding: utf-8 -*-
# from odoo import http


# class DisableMergeOnStockmove(http.Controller):
#     @http.route('/disable_merge_on_stockmove/disable_merge_on_stockmove', auth='public')
#     def index(self, **kw):
#         return "Hello, world"

#     @http.route('/disable_merge_on_stockmove/disable_merge_on_stockmove/objects', auth='public')
#     def list(self, **kw):
#         return http.request.render('disable_merge_on_stockmove.listing', {
#             'root': '/disable_merge_on_stockmove/disable_merge_on_stockmove',
#             'objects': http.request.env['disable_merge_on_stockmove.disable_merge_on_stockmove'].search([]),
#         })

#     @http.route('/disable_merge_on_stockmove/disable_merge_on_stockmove/objects/<model("disable_merge_on_stockmove.disable_merge_on_stockmove"):obj>', auth='public')
#     def object(self, obj, **kw):
#         return http.request.render('disable_merge_on_stockmove.object', {
#             'object': obj
#         })

