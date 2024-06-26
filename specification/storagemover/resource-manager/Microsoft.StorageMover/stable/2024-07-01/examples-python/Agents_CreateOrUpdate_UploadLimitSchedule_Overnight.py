from typing import Any, IO, Union

from azure.identity import DefaultAzureCredential

from azure.mgmt.storagemover import StorageMoverMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-storagemover
# USAGE
    python agents_create_or_update_upload_limit_schedule_overnight.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = StorageMoverMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="60bcfc77-6589-4da2-b7fd-f9ec9322cf95",
    )

    response = client.agents.create_or_update(
        resource_group_name="examples-rg",
        storage_mover_name="examples-storageMoverName",
        agent_name="examples-agentName",
        agent={
            "properties": {
                "arcResourceId": "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName",
                "arcVmUuid": "3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9",
                "uploadLimitSchedule": {
                    "weeklyRecurrences": [
                        {
                            "days": ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
                            "endTime": {"hour": 24, "minute": 0},
                            "limitInMbps": 2000,
                            "startTime": {"hour": 18, "minute": 0},
                        },
                        {
                            "days": ["Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"],
                            "endTime": {"hour": 9, "minute": 0},
                            "limitInMbps": 2000,
                            "startTime": {"hour": 0, "minute": 0},
                        },
                    ]
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/Agents_CreateOrUpdate_UploadLimitSchedule_Overnight.json
if __name__ == "__main__":
    main()
