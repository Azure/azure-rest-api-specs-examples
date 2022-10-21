
from azure.identity import DefaultAzureCredential
from azure.mgmt.cdn import CdnManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cdn
# USAGE
    python afd_origins_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret 
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CdnManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.afd_origins.begin_update(
        resource_group_name="RG",
        profile_name="profile1",
        origin_group_name="origingroup1",
        origin_name="origin1",
        origin_update_properties={
            "properties": {
                "enabledState": "Enabled",
                "hostName": "host1.blob.core.windows.net",
                "httpPort": 80,
                "httpsPort": 443,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/AFDOrigins_Update.json
if __name__ == "__main__":
    main()
