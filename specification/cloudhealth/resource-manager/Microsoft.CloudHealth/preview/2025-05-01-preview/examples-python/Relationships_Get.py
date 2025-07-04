from azure.identity import DefaultAzureCredential

from azure.mgmt.cloudhealth import CloudHealthMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-cloudhealth
# USAGE
    python relationships_get.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = CloudHealthMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.relationships.get(
        resource_group_name="rgopenapi",
        health_model_name="myHealthModel",
        relationship_name="Ue-21-F3M12V3w-13x18F8H-7HOk--kq6tP-HB",
    )
    print(response)


# x-ms-original-file: 2025-05-01-preview/Relationships_Get.json
if __name__ == "__main__":
    main()
