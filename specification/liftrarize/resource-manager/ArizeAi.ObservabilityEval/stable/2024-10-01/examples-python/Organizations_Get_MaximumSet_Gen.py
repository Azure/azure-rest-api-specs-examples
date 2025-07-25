from azure.identity import DefaultAzureCredential

from azure.mgmt.arizeaiobservabilityeval import ArizeAIObservabilityEvalMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-arizeaiobservabilityeval
# USAGE
    python organizations_get_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ArizeAIObservabilityEvalMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.organizations.get(
        resource_group_name="rgopenapi",
        organizationname="test-organization-1",
    )
    print(response)


# x-ms-original-file: 2024-10-01/Organizations_Get_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
