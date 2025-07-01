from odoo import models

class DisableMergeOnStockMove(models.Model):
    _inherit = 'stock.move'

    # def _action_confirm(self, merge=False):
    #     # Commented out the merging behavior
    #     # if merge:
    #     #     moves = self._merge_moves(merge_into=merge_into)
    #     #
    #     # Instead, proceed without merging
    #     moves = self

    #     # Call the original logic but skip the merging part
    #     moves.filtered(lambda move: move.state == 'draft')._set_quantity_done_zero()
    #     moves.filtered(lambda move: move.state == 'draft')._assign_picking()
    #     moves.write({'state': 'confirmed', 'date': fields.Datetime.now()})
    #     return moves

    def _action_confirm(self, merge=False):  # default changed to False
        return super()._action_confirm(merge=merge)
