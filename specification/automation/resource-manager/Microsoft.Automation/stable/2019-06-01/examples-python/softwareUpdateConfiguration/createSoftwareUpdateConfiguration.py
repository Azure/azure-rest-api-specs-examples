from azure.identity import DefaultAzureCredential

from azure.mgmt.automation import AutomationClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-automation
# USAGE
    python create_software_update_configuration.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AutomationClient(
        credential=DefaultAzureCredential(),
        subscription_id="51766542-3ed7-4a72-a187-0c8ab644ddab",
    )

    response = client.software_update_configurations.create(
        resource_group_name="mygroup",
        automation_account_name="myaccount",
        software_update_configuration_name="testpatch",
        parameters={
            "properties": {
                "scheduleInfo": {
                    "advancedSchedule": {"weekDays": ["Monday", "Thursday"]},
                    "expiryTime": "2018-11-09T11:22:57+00:00",
                    "frequency": "Hour",
                    "interval": 1,
                    "startTime": "2017-10-19T12:22:57+00:00",
                    "timeZone": "America/Los_Angeles",
                },
                "tasks": {
                    "postTask": {"parameters": None, "source": "GetCache"},
                    "preTask": {"parameters": {"COMPUTERNAME": "Computer1"}, "source": "HelloWorld"},
                },
                "updateConfiguration": {
                    "azureVirtualMachines": [
                        "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-01",
                        "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-02",
                        "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources/providers/Microsoft.Compute/virtualMachines/vm-03",
                    ],
                    "duration": "PT2H0M",
                    "nonAzureComputerNames": ["box1.contoso.com", "box2.contoso.com"],
                    "operatingSystem": "Windows",
                    "targets": {
                        "azureQueries": [
                            {
                                "locations": ["Japan East", "UK South"],
                                "scope": [
                                    "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067/resourceGroups/myresources",
                                    "/subscriptions/5ae68d89-69a4-454f-b5ce-e443cc4e0067",
                                ],
                                "tagSettings": {
                                    "filterOperator": "All",
                                    "tags": {
                                        "tag1": ["tag1Value1", "tag1Value2", "tag1Value3"],
                                        "tag2": ["tag2Value1", "tag2Value2", "tag2Value3"],
                                    },
                                },
                            }
                        ],
                        "nonAzureQueries": [
                            {"functionAlias": "SavedSearch1", "workspaceId": "WorkspaceId1"},
                            {"functionAlias": "SavedSearch2", "workspaceId": "WorkspaceId2"},
                        ],
                    },
                    "windows": {
                        "excludedKbNumbers": ["168934", "168973"],
                        "includedUpdateClassifications": "Critical",
                        "rebootSetting": "IfRequired",
                    },
                },
            }
        },
    )
    print(response)


# x-ms-original-file: specification/automation/resource-manager/Microsoft.Automation/stable/2019-06-01/examples/softwareUpdateConfiguration/createSoftwareUpdateConfiguration.json
if __name__ == "__main__":
    main()
