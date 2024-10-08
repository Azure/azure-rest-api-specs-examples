from azure.identity import DefaultAzureCredential

from azure.mgmt.azurestackhci import AzureStackHCIClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-azurestackhci
# USAGE
    python put_update_runs.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AzureStackHCIClient(
        credential=DefaultAzureCredential(),
        subscription_id="b8d594e5-51f3-4c11-9c54-a7771b81c712",
    )

    response = client.update_runs.put(
        resource_group_name="testrg",
        cluster_name="testcluster",
        update_name="Microsoft4.2203.2.32",
        update_run_name="23b779ba-0d52-4a80-8571-45ca74664ec3",
        update_runs_properties={
            "properties": {
                "progress": {
                    "description": "Update Azure Stack.",
                    "endTimeUtc": "2022-04-06T13:58:42.969006+00:00",
                    "errorMessage": "",
                    "lastUpdatedTimeUtc": "2022-04-06T13:58:42.969006+00:00",
                    "name": "Unnamed step",
                    "startTimeUtc": "2022-04-06T01:36:33.3876751+00:00",
                    "status": "Success",
                    "steps": [
                        {
                            "description": "Prepare for SSU update",
                            "endTimeUtc": "2022-04-06T01:37:16.8728314+00:00",
                            "errorMessage": "",
                            "lastUpdatedTimeUtc": "2022-04-06T01:37:16.8728314+00:00",
                            "name": "PreUpdate Cloud",
                            "startTimeUtc": "2022-04-06T01:36:33.3876751+00:00",
                            "status": "Success",
                            "steps": [],
                        }
                    ],
                }
            }
        },
    )
    print(response)


# x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/PutUpdateRuns.json
if __name__ == "__main__":
    main()
