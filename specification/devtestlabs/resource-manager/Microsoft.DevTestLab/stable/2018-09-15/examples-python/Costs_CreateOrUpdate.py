from azure.identity import DefaultAzureCredential

from azure.mgmt.devtestlabs import DevTestLabsClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-devtestlabs
# USAGE
    python costs_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = DevTestLabsClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscriptionId}",
    )

    response = client.costs.create_or_update(
        resource_group_name="resourceGroupName",
        lab_name="{labName}",
        name="targetCost",
        lab_cost={
            "properties": {
                "currencyCode": "USD",
                "endDateTime": "2020-12-31T23:59:59Z",
                "startDateTime": "2020-12-01T00:00:00Z",
                "targetCost": {
                    "costThresholds": [
                        {
                            "displayOnChart": "Disabled",
                            "percentageThreshold": {"thresholdValue": 25},
                            "sendNotificationWhenExceeded": "Disabled",
                            "thresholdId": "00000000-0000-0000-0000-000000000001",
                        },
                        {
                            "displayOnChart": "Enabled",
                            "percentageThreshold": {"thresholdValue": 50},
                            "sendNotificationWhenExceeded": "Enabled",
                            "thresholdId": "00000000-0000-0000-0000-000000000002",
                        },
                        {
                            "displayOnChart": "Disabled",
                            "percentageThreshold": {"thresholdValue": 75},
                            "sendNotificationWhenExceeded": "Disabled",
                            "thresholdId": "00000000-0000-0000-0000-000000000003",
                        },
                        {
                            "displayOnChart": "Disabled",
                            "percentageThreshold": {"thresholdValue": 100},
                            "sendNotificationWhenExceeded": "Disabled",
                            "thresholdId": "00000000-0000-0000-0000-000000000004",
                        },
                        {
                            "displayOnChart": "Disabled",
                            "percentageThreshold": {"thresholdValue": 125},
                            "sendNotificationWhenExceeded": "Disabled",
                            "thresholdId": "00000000-0000-0000-0000-000000000005",
                        },
                    ],
                    "cycleEndDateTime": "2020-12-31T00:00:00.000Z",
                    "cycleStartDateTime": "2020-12-01T00:00:00.000Z",
                    "cycleType": "CalendarMonth",
                    "status": "Enabled",
                    "target": 100,
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Costs_CreateOrUpdate.json
if __name__ == "__main__":
    main()
