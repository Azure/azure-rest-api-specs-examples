from azure.identity import DefaultAzureCredential
from azure.mgmt.securityinsight import SecurityInsights

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-securityinsight
# USAGE
    python get_metadata.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SecurityInsights(
        credential=DefaultAzureCredential(),
        subscription_id="2e1dc338-d04d-4443-b721-037eff4fdcac",
    )

    response = client.metadata.get(
        resource_group_name="myRg",
        workspace_name="myWorkspace",
        metadata_name="metadataName",
    )
    print(response)


# x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-12-01-preview/examples/metadata/GetMetadata.json
if __name__ == "__main__":
    main()
