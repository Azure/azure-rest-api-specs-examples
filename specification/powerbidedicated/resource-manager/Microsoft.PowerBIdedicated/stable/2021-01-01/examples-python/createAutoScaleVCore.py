from azure.identity import DefaultAzureCredential
from azure.mgmt.powerbidedicated import PowerBIDedicated

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-powerbidedicated
# USAGE
    python create_auto_scale_vcore.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = PowerBIDedicated(
        credential=DefaultAzureCredential(),
        subscription_id="613192d7-503f-477a-9cfe-4efc3ee2bd60",
    )

    response = client.auto_scale_vcores.create(
        resource_group_name="TestRG",
        vcore_name="testvcore",
        v_core_parameters={
            "location": "West US",
            "properties": {"capacityLimit": 10, "capacityObjectId": "a28f00bd-5330-4572-88f1-fa883e074785"},
            "sku": {"capacity": 0, "name": "AutoScale", "tier": "AutoScale"},
            "tags": {"testKey": "testValue"},
        },
    )
    print(response)


# x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/createAutoScaleVCore.json
if __name__ == "__main__":
    main()
