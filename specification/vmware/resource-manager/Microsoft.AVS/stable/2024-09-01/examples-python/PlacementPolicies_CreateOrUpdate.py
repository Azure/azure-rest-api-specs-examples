from azure.identity import DefaultAzureCredential

from azure.mgmt.avs import AVSClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-avs
# USAGE
    python placement_policies_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = AVSClient(
        credential=DefaultAzureCredential(),
        subscription_id="00000000-0000-0000-0000-000000000000",
    )

    response = client.placement_policies.begin_create_or_update(
        resource_group_name="group1",
        private_cloud_name="cloud1",
        cluster_name="cluster1",
        placement_policy_name="policy1",
        placement_policy={
            "properties": {
                "affinityStrength": "Must",
                "affinityType": "AntiAffinity",
                "azureHybridBenefitType": "SqlHost",
                "hostMembers": [
                    "fakehost22.nyc1.kubernetes.center",
                    "fakehost23.nyc1.kubernetes.center",
                    "fakehost24.nyc1.kubernetes.center",
                ],
                "type": "VmHost",
                "vmMembers": [
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-128",
                    "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/group1/providers/Microsoft.AVS/privateClouds/cloud1/clusters/cluster1/virtualMachines/vm-256",
                ],
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2024-09-01/examples/PlacementPolicies_CreateOrUpdate.json
if __name__ == "__main__":
    main()
