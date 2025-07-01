{
    'name': 'Disable Merge on Stock Move',
    'version': '1.0',
    'summary': 'Prevents automatic merging of stock move lines',
    'description': """
        This module disables the default behavior that merges stock move lines under certain conditions.
    """,
    'category': 'Inventory',
    'author': 'Your Company Name',
    'website': 'https://www.example.com',
    'depends': ['base','stock'], 
    'data': [
        # Add your XML/CSV files here if any
    ],
    'installable': True,
    'application': False,
    'auto_install': False,
    'license': 'LGPL-3',
}
