from azure.identity import DefaultAzureCredential
from azure.mgmt.subscription import SubscriptionClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-subscription
# USAGE
    python create_alias.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = SubscriptionClient(
        credential=DefaultAzureCredential(),
    )

    response = client.alias.begin_create(
        alias_name="aliasForNewSub",
        body={
            "properties": {
                "additionalProperties": {
                    "managementGroupId": None,
                    "subscriptionOwnerId": "f09b39eb-c496-482c-9ab9-afd799572f4c",
                    "subscriptionTenantId": "66f6e4d6-07dc-4aea-94ea-e12d3026a3c8",
                    "tags": {"tag1": "Messi", "tag2": "Ronaldo", "tag3": "Lebron"},
                },
                "billingScope": "/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH",
                "displayName": "Test Subscription",
                "subscriptionId": None,
                "workload": "Production",
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/subscription/resource-manager/Microsoft.Subscription/stable/2021-10-01/examples/createAlias.json
if __name__ == "__main__":
    main()
