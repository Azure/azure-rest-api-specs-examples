from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python addons_create_or_update_srm.py

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
        addon_name="srm",
        addon={"properties": {"addonType": "SRM", "licenseKey": "41915178-A8FF-4A4D-B683-6D735AF5E3F5"}},
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/Addons_CreateOrUpdate_SRM.json
if __name__ == "__main__":
    main()
