from azure.identity import DefaultAzureCredential
from azure.mgmt.advisor import AdvisorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-advisor
# USAGE
    python create_suppression.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AdvisorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.suppressions.create(
        resource_uri="resourceUri",
        recommendation_id="recommendationId",
        name="suppressionName1",
        suppression_contract={"properties": {"ttl": "07:00:00:00"}},
    )
    print(response)


# x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateSuppression.json
if __name__ == "__main__":
    main()
