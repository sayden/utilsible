# WIP Utilsible

Command line utilities for Ansible

## Commands

### Add Role `utilsible roles add [role name]`
```bash
    utilsible roles add my_role
```

If no roles folder is found and it isn't executed in the root folder, creates a "roles" folder and a "my_role" folder within it.

Inside the "my_role" folder, it will create:

* `tasks/main.yml` with the content `# roles/my_role/tasks/main.yml`
* `templates/main.yml`  with the content `# roles/my_role/templates/main.yml`
* `vars/main.yml` etc...
* `files/main.yml`
* `meta/main.yml`
* `handlers/main.yml`
* `defaults/main.yml`
