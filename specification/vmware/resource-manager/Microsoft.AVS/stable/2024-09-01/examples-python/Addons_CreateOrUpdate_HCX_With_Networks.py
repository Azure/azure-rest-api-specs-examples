from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python addons_create_or_update_hcx_with_networks.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AVSClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.addons.begin_create_or_update(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        addon_name="hcx",
        addon={
            "properties": {
                "addonType": "HCX",
                "managementNetwork": "10.3.1.0/24",
                "offer": "VMware MaaS Cloud Provider (Enterprise)",
                "uplinkNetwork": "10.3.2.0/24",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/Addons_CreateOrUpdate_HCX_With_Networks.json
if __name__ == "__main__":
    main()
