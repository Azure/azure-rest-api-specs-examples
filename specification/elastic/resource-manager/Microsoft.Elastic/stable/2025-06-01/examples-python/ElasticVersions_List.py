from azure.identity import DefaultAzureCredential

from azure.mgmt.elastic import ElasticMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-elastic
# USAGE
    python elastic_versions_list.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ElasticMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.elastic_versions.list(
        region="myregion",
    )
    for item in response:
        print(item)


# x-ms-original-file: 2025-06-01/ElasticVersions_List.json
if __name__ == "__main__":
    main()
