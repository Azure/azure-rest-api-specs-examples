from azure.identity import DefaultAzureCredential

from azure.mgmt.hardwaresecuritymodules import HardwareSecurityModulesMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hardwaresecuritymodules
# USAGE
    python dedicated_hsm_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HardwareSecurityModulesMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.dedicated_hsm.begin_create_or_update(
        resource_group_name="hsm-group",
        name="hsm1",
        parameters={
            "location": "westus",
            "properties": {
                "networkProfile": {
                    "networkInterfaces": [{"privateIpAddress": "1.0.0.1"}],
                    "subnet": {
                        "resourceId": "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/hsm-group/providers/Microsoft.Network/virtualNetworks/stamp01/subnets/stamp01"
                    },
                },
                "stampId": "stamp01",
            },
            "sku": {"name": "SafeNet Luna Network HSM A790"},
            "tags": {"Dept": "hsm", "Environment": "dogfood"},
        },
    ).result()
    print(response)


# x-ms-original-file: 2025-03-31/DedicatedHsm_CreateOrUpdate.json
if __name__ == "__main__":
    main()
