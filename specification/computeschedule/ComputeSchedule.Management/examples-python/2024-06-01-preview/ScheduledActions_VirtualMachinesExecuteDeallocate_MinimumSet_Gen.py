from azure.identity import DefaultAzureCredential

from azure.mgmt.computeschedule import ComputeScheduleMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-computeschedule
# USAGE
    python scheduled_actions_virtual_machines_execute_deallocate_minimum_set_gen.py

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

    response = client.scheduled_actions.virtual_machines_execute_deallocate(
        locationparameter="irbtacrhjgfpyvhvjo",
        request_body={
            "correlationid": "01080d2f-1dca-4610-afb4-dd25eddf1f3c1",
            "executionParameters": {},
            "resources": {
                "ids": [
                    "/subscriptions/YourSubscriptionId/resourceGroups/YourResourceGroupName/providers/Microsoft.Compute/virtualMachines/testResource2"
                ]
            },
        },
    )
    print(response)


# x-ms-original-file: 2024-08-15-preview/ScheduledActions_VirtualMachinesExecuteDeallocate_MinimumSet_Gen.json
if __name__ == "__main__":
    main()
