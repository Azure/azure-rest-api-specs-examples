from azure.identity import DefaultAzureCredential

from azure.mgmt.netapp import NetAppManagementClient

"""
# PREREQUISITES
    pip install azure-identity
    pip install azure-mgmt-netapp
# USAGE
    python buckets_create_or_update.py

    Before run the sample, please set the values of the client ID, tenant ID and client secret
    of the AAD application as environment variables: AZURE_CLIENT_ID, AZURE_TENANT_ID,
    AZURE_CLIENT_SECRET. For more info about how to get the value, please see:
    https://docs.microsoft.com/azure/active-directory/develop/howto-create-service-principal-portal
"""


def main():
    client = NetAppManagementClient(
        credential=DefaultAzureCredential(),
        subscription_id="D633CC2E-722B-4AE1-B636-BBD9E4C60ED9",
    )

    response = client.buckets.begin_create_or_update(
        resource_group_name="myRG",
        account_name="account1",
        pool_name="pool1",
        volume_name="volume1",
        bucket_name="bucket1",
        body={
            "properties": {
                "fileSystemUser": {"nfsUser": {"groupId": 1000, "userId": 1001}},
                "path": "/path",
                "server": {
                    "certificateObject": "LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSURuVENDQW9XZ0F3SUJBZ0lRR3FXdnRxeHBvSTFJV3Z4VGdJbElWREFOQmdrcWhraUc5dzBCQVFzRkFEQlYKTVJNd0VRWUtDWkltaVpQeUxHUUJHUllEWTI5dE1SY3dGUVlLQ1pJbWlaUHlMR1FCR1JZSGFHRnlhV3R5WWpFbApNQ01HQTFVRUF4TWNhR0Z5YVd0eVlpMVhTVTR0TWtKUFZrRkZTMEkwTkVJdFEwRXRNakFlRncweU1EQTFNRFV3Ck56TTVORGxhRncweU1EQTFNRFl3TnpRNU5EaGFNRlV4RXpBUkJnb0praWFKay9Jc1pBRVpGZ05qYjIweEZ6QVYKQmdvSmtpYUprL0lzWkFFWkZnZG9ZWEpwYTNKaU1TVXdJd1lEVlFRREV4eG9ZWEpwYTNKaUxWZEpUaTB5UWs5VwpRVVZMUWpRMFFpMURRUzB5TUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUFqMHByCnhaaXpNaDBqYnRwN1ZOc0JrRVJ2MVpZT0MzMEtqaGRWdEExRm1MeFM2cXlycmpMZUdXOXRSd2ZnUkR0eVBodTIKZVJTcVpTUjF6Z1hZR0s0Nys3Y3F0YnB2UElOektCb0dOWERIVTNxVWlleXJWSjFDVzRKNjJodUdrbUV1VVVkMApKMXBxNTVxbjk1SmRUbWh1dmZlTUxxeHB5c01nbGVnY281ZFhoN0hsQkhwaTNKMFN4ZnhVWmxKMVZiOFJZVEZhCkJiMGFlTVZaRzRKeVREaktiMlR1TmFXOG1aUE5vOFBMRDRocjdndFNZUEQvQ1dVVGV5QlpoZC9LTzNPczlWVEIKYmpLUGtWd0J2WEs2SlFMSGprNFBHS3VYZDhaWVFyajBtOWNIZDNmcWNYTXlQUnQ2TlJ4ak0yMTUxckFzSkVhNgpWZC9ta056akpXalBrT2VZUVFJREFRQUJvMmt3WnpBVEJna3JCZ0VFQVlJM0ZBSUVCaDRFQUVNQVFUQU9CZ05WCkhROEJBZjhFQkFNQ0FZWXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVVDVEc2ODJSay9kMysKWGtHa0VMakRFMjI4ZjNnd0VBWUpLd1lCQkFHQ054VUJCQU1DQVFBd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQgpBR1RjTTNnWExIU05wS014RHEvUFpZbWZCMmNlN3VhMmxxbXpzZSt4QmthSTE0WXdzZE5ZQjZBVTNFWDM3QWpZCjd3bm9xQzY1N0U2RVhTODVDckoyNXJNTHo4OEtONGI3cUg5RUowSS9XVHg5YTdUT0ZENENWQThuL0xwNGh1Ym4KNlBFalY5NFlZWXBXTG1hTkkvbGFReWsxSHVJbDFSTCttVDFnSWQ4ZWZXZ1UvNmlVVEw3eGMrdjkyNHBuTHhISwpOSnNTV3c0NFk5a0R5SU9KOXFjWUlBN1lhTkxPZTRjSysvQlRvdDh0dVVKT1hHLzdBRmtxR2EyQVA4MmFZOStKCnkwSmU2OG5nTHJ1dVU4VHpneVpqdkFHcTRrVEVOdWFoaFdHVC9KWkEzOXhSNUV4MmNMUUplcE5NdnlZbUZ3Z1UKME8zYlA0OWNBVFVCMXoyQ3Y5aTRQbVk9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K",
                    "fqdn": "fullyqualified.domainname.com",
                },
            }
        },
    ).result()
    print(response)


# x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/Buckets_CreateOrUpdate.json
if __name__ == "__main__":
    main()
