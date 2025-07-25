from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.policy import PolicyClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python create_policy_assignment_by_id.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PolicyClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.policy_assignments.create_by_id(
        policy_assignment_id="providers/Microsoft.Management/managementGroups/MyManagementGroup/providers/Microsoft.Authorization/policyAssignments/LowCostStorage",
        parameters={
            "properties": {
                "definitionVersion": "1.*.*",
                "description": "Allow only storage accounts of SKU Standard_GRS or Standard_LRS to be created",
                "displayName": "Enforce storage account SKU",
                "enforcementMode": "Default",
                "metadata": {"assignedBy": "Cheapskate Boss"},
                "parameters": {"listOfAllowedSKUs": {"value": ["Standard_GRS", "Standard_LRS"]}},
                "policyDefinitionId": "/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1",
            }
        },
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/stable/2023-04-01/examples/createPolicyAssignmentById.json
if __name__ == "__main__":
    main()
