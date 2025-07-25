from azure.identity import DefaultAzureCredential

from azure.mgmt.resource.policy import PolicyClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-resource
# USAGE
    python create_or_update_policy_exemption_with_resource_selectors.py

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

    response = client.policy_exemptions.create_or_update(
        scope="subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster",
        policy_exemption_name="DemoExpensiveVM",
        parameters={
            "properties": {
                "assignmentScopeValidation": "Default",
                "description": "Exempt demo cluster from limit sku",
                "displayName": "Exempt demo cluster",
                "exemptionCategory": "Waiver",
                "metadata": {"reason": "Temporary exemption for a expensive VM demo"},
                "policyAssignmentId": "/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement",
                "policyDefinitionReferenceIds": ["Limit_Skus"],
                "resourceSelectors": [
                    {
                        "name": "SDPRegions",
                        "selectors": [{"in": ["eastus2euap", "centraluseuap"], "kind": "resourceLocation"}],
                    }
                ],
            }
        },
    )
    print(response)


# x-ms-original-file: specification/resources/resource-manager/Microsoft.Authorization/preview/2022-07-01-preview/examples/createOrUpdatePolicyExemptionWithResourceSelectors.json
if __name__ == "__main__":
    main()
