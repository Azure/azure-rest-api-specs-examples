
import com.azure.resourcemanager.recoveryservices.models.AuthType;
import com.azure.resourcemanager.recoveryservices.models.RawCertificateData;

/**
 * Samples for VaultCertificates Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2025-02-01/examples/
     * PUTVaultCred.json
     */
    /**
     * Sample code: Download vault credential file.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        downloadVaultCredentialFile(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.vaultCertificates()
            .define("BCDRIbzVault77777777-d41f-4550-9f70-7708a3a2283b-12-18-2017-vaultcredentials")
            .withExistingVault("BCDRIbzRG", "BCDRIbzVault")
            .withProperties(new RawCertificateData().withAuthType(AuthType.AAD).withCertificate(
                "TUlJRE5EQ0NBaHlnQXdJQkFnSVFDYUxFKzVTSlNVeWdncDM0VS9HUm9qQU5CZ2txaGtpRzl3MEJBUXNGQURBWE1SVXdFd1lEVlFRREV3eGhiV05vWVc1a2JpNWpiMjB3SGhjTk1qSXhNREkwTVRJd05qRTRXaGNOTWpNeE1ESTBNVEl4TmpFNFdqQVhNUlV3RXdZRFZRUURFd3hoYldOb1lXNWtiaTVqYjIwd2dnRWlNQTBHQ1NxR1NJYjNEUUVCQVFVQUE0SUJEd0F3Z2dFS0FvSUJBUUN4cFpwS293a2p4VU9VWkpLT2JvdGdPWXkzaW9UVkxMMmZyaW9nZVN1Qm5IMWw3aVdQWW9kUHRoWS8yVmh6ZFVUckNXL25pNUh3b0JHYzZMMHF6UGlBWXpHek94RmpMQjZjdFNkbm9nL1A4eEV2OGE0cnJWZlBZdS9INStoTGx3N0RubXlTNWs4TU9sSVhUemVWNkxZV2I2RWlpTFppc0k1R3lLU1liemNaQmJKdnhLTVdGdHRCV08xZUwzUWNUejlpb1VGQzVnRlFKQzg3YXFkeDR1Wk9WYzRLM3Ixb09sTFBKdmRLN25YU3VWci9ZOC80ZHhCdDJZUTRia0hjM2EzcUNBbTZrV0QzamRiajhCZmhlWWNVNjFFZ3llVFV2MlI4dzRubWJqVXZxRW05cDZtTG4xMTdEWWpQTHNFODVTL0FpQmF0dkNhQ3hCZ0lxb1N1blBOUkFnTUJBQUdqZkRCNk1BNEdBMVVkRHdFQi93UUVBd0lGb0RBSkJnTlZIUk1FQWpBQU1CMEdBMVVkSlFRV01CUUdDQ3NHQVFVRkJ3TUJCZ2dyQmdFRkJRY0RBakFmQmdOVkhTTUVHREFXZ0JRR1NZcDJMUTJwOE5wMHUzRThJZDdRUjRTQXBqQWRCZ05WSFE0RUZnUVVCa21LZGkwTnFmRGFkTHR4UENIZTBFZUVnS1l3RFFZSktvWklodmNOQVFFTEJRQURnZ0VCQUp2ZG9yRmJ4cExZaUhYRHpnR001WmxMWTRDZE1LYW5BdzVDZDNFVnhDbkhtT05ISnpLRmpzdHZjdUN1TDZ2S1ptci9abm5ENXNLUnE0d0xnTXV6dlNXNGtQTXlWeENrYzdVYnNZSWJCSXNIUDl3cUNmcUY5aG5LSE9YZFJJV2tBVXhnbmYxSlpLZjR1NlpTSzZ3dExaME9VT0c5Mmd3SlB2eW5PVmJoeWpqczdQTVpONEw1djZyeHJkRWp0WG5sYzIvRDlnS0NOTFhFZHdRM0dzS05ZTGZvYy9DT3JmbEIrRHVPSThrVzM0WmxzYlFHelgyQ3ArWVVlSDNrQlBjY3RpUWNURHFQcW5YS0NNMTJ6MGZDTjVpNXRkRlUrM0VzemZBQkpiOEZpU2ZCWFF1UUZRRDNDTDkraVdjZXhrMmxQako2akZIbHZtak9XbTdjQllHZlc4ST0="
                    .getBytes()))
            .create();
    }
}
