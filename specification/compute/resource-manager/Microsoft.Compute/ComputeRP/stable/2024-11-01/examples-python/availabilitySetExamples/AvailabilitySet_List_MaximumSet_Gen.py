from azure.identity import DefaultAzureCredential

from azure.mgmt.compute import ComputeManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-compute
# USAGE
    python availability_set_list_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ComputeManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="{subscription-id}",
    )

    response = client.availability_sets.list(
        resource_group_name="rgcompute",
    )
    for item in response:
        print(item)


# x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-11-01/examples/availabilitySetExamples/AvailabilitySet_List_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
