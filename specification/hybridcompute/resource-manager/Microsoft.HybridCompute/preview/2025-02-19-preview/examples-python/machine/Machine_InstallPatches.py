from azure.identity import DefaultAzureCredential

from azure.mgmt.hybridcompute import HybridComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-hybridcompute
# USAGE
    python machine_install_patches.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HybridComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.machines.begin_install_patches(
        resource_group_name="myResourceGroupName",
        name="myMachineName",
        install_patches_input={
            "maximumDuration": "PT4H",
            "rebootSetting": "IfRequired",
            "windowsParameters": {
                "classificationsToInclude": ["Critical", "Security"],
                "maxPatchPublishDate": "2021-08-19T02:36:43.0539904+00:00",
                "patchNameMasksToExclude": ["*Windows*"],
                "patchNameMasksToInclude": ["*SQL*"],
            },
        },
    ).result()
    print(response)


# x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/machine/Machine_InstallPatches.json
if __name__ == "__main__":
    main()
