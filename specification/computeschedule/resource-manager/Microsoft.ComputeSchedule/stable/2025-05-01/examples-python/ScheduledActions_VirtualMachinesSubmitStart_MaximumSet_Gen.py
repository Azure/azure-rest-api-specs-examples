from azure.identity import DefaultAzureCredential

from azure.mgmt.computeschedule import ComputeScheduleMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-computeschedule
# USAGE
    python scheduled_actions_virtual_machines_submit_start_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeScheduleMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.scheduled_actions.virtual_machines_submit_start(
        locationparameter="bgyvpodcjmcmbxohvil",
        request_body={
            "correlationid": "bvmpxvbd",
            "executionParameters": {
                "optimizationPreference": "Cost",
                "retryPolicy": {"retryCount": 25, "retryWindowInMinutes": 4},
            },
            "resources": {
                "ids": [
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource3"
                ]
            },
            "schedule": {
                "deadLine": "2025-04-15T19:47:04.403Z",
                "deadline": "2025-04-15T19:47:04.403Z",
                "deadlineType": "Unknown",
                "timeZone": "upnmayfebiadztdktxzq",
                "timezone": "qacufsmctpgjozovlsihrzoctatcsj",
            },
        },
    )
    print(response)


# x-ms-original-file: 2025-05-01/ScheduledActions_VirtualMachinesSubmitStart_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
