# Node management
See https://tailscale.com/kb/1099/device-approval for more information.

## Setup
### 1. Change the configuration
1. Change the `config.yaml` to contain the desired records like so:

    ```yaml
    node_management:
      manual_approve_new_node: true
    ```

2. Restart your headscale instance.

## Usage

### Pre-approve a node using preauthkeys
1. Create preauthkey with a pre-approve option
```bash
headscale preauthkeys create --user=<USER_NAME> --pre-approved  
```

2. Register a node on the headscale using preauthkey (with the pre-approval option enabled)
```bash
headscale nodes register --user=<USER_NAME> --mkey=mkey:<MACHINE_KEY> --auth-key=<PREAUTHKEY_PRE_APPROVED>
```

### Approve a node after registration without the pre-authkey pre-approval option
```bash
headscale nodes approve --identifier=<NODE_ID>
```
