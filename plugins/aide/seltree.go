package aide

import (
  "container/list"
)

/* seltree structure
 * lists have regex_t* in them
 * checked is whether or not the node has been checked yet and status
 * when added
 * path is the path of the node
 * parent is the parent, NULL if root
 * childs is list of seltree*:s
 * new_data is this nodes new attributes (read from disk or db in --compare)
 * old_data is this nodes old attributes (read from db)
 * attr attributes to add for this node and possibly for its children
 * changed_attrs changed attributes between new_data and old_data
 */

type Seltree struct {

	Sel_rx_lst *list.List
	Neg_rx_lst *list.List
	Equ_rx_lst *list.List
	Childs *list.List
	Parent *Seltree

	Path string
	Checked uint64

	Conf_lineno int64
	Rx string

	Attr uint64

	New_data *DB_Line
	Old_data *DB_Line

	Changed_attrs uint64

}

