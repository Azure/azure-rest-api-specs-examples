from azure.identity import DefaultAzureCredential

from azure.mgmt.neonpostgres import NeonPostgresMgmtClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-neonpostgres
# USAGE
    python projects_get_connection_uri_maximum_set_gen.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NeonPostgresMgmtClient(
        credential=DefaultAzureCredential(),
        subscription_id="SUBSCRIPTION_ID",
    )

    response = client.projects.get_connection_uri(
        resource_group_name="rgneon",
        organization_name="test-org",
        project_name="entity-name",
        connection_uri_parameters={
            "branchId": "iimmlbqv",
            "databaseName": "xc",
            "endpointId": "jcpdvsyjcn",
            "isPooled": True,
            "projectId": "riuifmoqtorrcffgksvfcobia",
            "roleName": "xhmcvsgtp",
        },
    )
    print(response)


# x-ms-original-file: 2025-03-01/Projects_GetConnectionUri_MaximumSet_Gen.json
if __name__ == "__main__":
    main()
