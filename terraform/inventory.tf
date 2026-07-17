resource "local_file" "ansible_inventory" {

  filename = "${path.module}/../ansible/inventory/hosts.ini"

  content = <<EOF
[jenkins]
jenkins-server ansible_host=${aws_instance.jenkins.public_ip}

[kubernetes_master]
master ansible_host=${aws_instance.master.public_ip}

[kubernetes_workers]
worker1 ansible_host=${aws_instance.worker1.public_ip}
worker2 ansible_host=${aws_instance.worker2.public_ip}

[kubernetes:children]
kubernetes_master
kubernetes_workers

[all:vars]
ansible_user=ubuntu
ansible_ssh_private_key_file=/root/.ssh/project.pem
ansible_python_interpreter=/usr/bin/python3
EOF

}
