from azure.identity import DefaultAzureCredential
from azure.mgmt.apicenter import ApiCenterMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-apicenter
# USAGE
    python api_versions_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ApiCenterMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.api_versions.get(
        resource_group_name="contoso-resources",
        service_name="contoso",
        workspace_name="default",
        api_name="echo-api",
        version_name="2023-01-01",
    )
    print(response)


# x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/stable/2024-03-01/examples/ApiVersions_Get.json
if __name__ == "__main__":
    main()
