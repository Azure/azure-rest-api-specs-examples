from azure.identity import DefaultAzureCredential

from azure.mgmt.managedservices import ManagedServicesClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-managedservices
# USAGE
    python delete_registration_assignment.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = ManagedServicesClient(
        credential=DefaultAzureCredential(),
    )

    client.registration_assignments.begin_delete(
        scope="subscription/0afefe50-734e-4610-8a82-a144ahf49dea",
        registration_assignment_id="26c128c2-fefa-4340-9bb1-6e081c90ada2",
    ).result()


# x-ms-original-file: specification/managedservices/resource-manager/Microsoft.ManagedServices/stable/2022-10-01/examples/DeleteRegistrationAssignment.json
if __name__ == "__main__":
    main()
