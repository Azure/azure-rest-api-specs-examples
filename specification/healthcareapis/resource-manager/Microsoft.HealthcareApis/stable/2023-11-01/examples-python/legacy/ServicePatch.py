from azure.identity import DefaultAzureCredential
from azure.mgmt.healthcareapis import HealthcareApisManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-healthcareapis
# USAGE
    python service_patch.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = HealthcareApisManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="subid",
    )

    response = client.services.begin_update(
        resource_group_name="rg1",
        resource_name="service1",
        service_patch_description={"tags": {"tag1": "value1", "tag2": "value2"}},
    ).result()
    print(response)


# x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2023-11-01/examples/legacy/ServicePatch.json
if __name__ == "__main__":
    main()
