from azure.identity import DefaultAzureCredential
from azure.mgmt.advisor import AdvisorManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-advisor
# USAGE
    python put_configurations.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AdvisorManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subscriptionId",
    )

    response = client.configurations.create_in_subscription(
        configuration_name="default",
        config_contract={
            "properties": {
                "digests": [
                    {
                        "actionGroupResourceId": "/subscriptions/subscriptionId/resourceGroups/resourceGroup/providers/microsoft.insights/actionGroups/actionGroupName",
                        "categories": ["HighAvailability", "Security", "Performance", "Cost", "OperationalExcellence"],
                        "frequency": 30,
                        "language": "en",
                        "name": "digestConfigName",
                        "state": "Active",
                    }
                ],
                "exclude": True,
                "lowCpuThreshold": "5",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/advisor/resource-manager/Microsoft.Advisor/stable/2020-01-01/examples/CreateConfiguration.json
if __name__ == "__main__":
    main()
