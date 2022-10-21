
from azure.identity import DefaultAzureCredential
from azure.mgmt.cdn import CdnManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cdn
# USAGE
    python origins_update.py

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

    response = client.origins.begin_update(
        resource_group_name="RG",
        profile_name="profile1",
        endpoint_name="endpoint1",
        origin_name="www-someDomain-net",
        origin_update_properties={
            "properties": {
                "enabled": True,
                "httpPort": 42,
                "httpsPort": 43,
                "originHostHeader": "www.someDomain2.net",
                "priority": 1,
                "privateLinkAlias": "APPSERVER.d84e61f0-0870-4d24-9746-7438fa0019d1.westus2.azure.privatelinkservice",
                "weight": 50,
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Origins_Update.json
if __name__ == "__main__":
    main()
