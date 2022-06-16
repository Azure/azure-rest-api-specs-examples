import com.azure.resourcemanager.confidentialledger.models.AadBasedSecurityPrincipal;
import com.azure.resourcemanager.confidentialledger.models.CertBasedSecurityPrincipal;
import com.azure.resourcemanager.confidentialledger.models.LedgerProperties;
import com.azure.resourcemanager.confidentialledger.models.LedgerRoleName;
import com.azure.resourcemanager.confidentialledger.models.LedgerType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/** Samples for Ledger Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/stable/2022-05-13/examples/ConfidentialLedger_Create.json
     */
    /**
     * Sample code: ConfidentialLedgerCreate.
     *
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void confidentialLedgerCreate(
        com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        manager
            .ledgers()
            .define("DummyLedgerName")
            .withExistingResourceGroup("DummyResourceGroupName")
            .withRegion("EastUS")
            .withTags(mapOf("additionalProps1", "additional properties"))
            .withProperties(
                new LedgerProperties()
                    .withLedgerType(LedgerType.PUBLIC)
                    .withAadBasedSecurityPrincipals(
                        Arrays
                            .asList(
                                new AadBasedSecurityPrincipal()
                                    .withPrincipalId("34621747-6fc8-4771-a2eb-72f31c461f2e")
                                    .withTenantId("bce123b9-2b7b-4975-8360-5ca0b9b1cd08")
                                    .withLedgerRoleName(LedgerRoleName.ADMINISTRATOR)))
                    .withCertBasedSecurityPrincipals(
                        Arrays
                            .asList(
                                new CertBasedSecurityPrincipal()
                                    .withCert(
                                        "-----BEGIN"
                                            + " CERTIFICATE-----MIIBsjCCATigAwIBAgIUZWIbyG79TniQLd2UxJuU74tqrKcwCgYIKoZIzj0EAwMwEDEOMAwGA1UEAwwFdXNlcjAwHhcNMjEwMzE2MTgwNjExWhcNMjIwMzE2MTgwNjExWjAQMQ4wDAYDVQQDDAV1c2VyMDB2MBAGByqGSM49AgEGBSuBBAAiA2IABBiWSo/j8EFit7aUMm5lF+lUmCu+IgfnpFD+7QMgLKtxRJ3aGSqgS/GpqcYVGddnODtSarNE/HyGKUFUolLPQ5ybHcouUk0kyfA7XMeSoUA4lBz63Wha8wmXo+NdBRo39qNTMFEwHQYDVR0OBBYEFPtuhrwgGjDFHeUUT4nGsXaZn69KMB8GA1UdIwQYMBaAFPtuhrwgGjDFHeUUT4nGsXaZn69KMA8GA1UdEwEB/wQFMAMBAf8wCgYIKoZIzj0EAwMDaAAwZQIxAOnozm2CyqRwSSQLls5r+mUHRGRyXHXwYtM4Dcst/VEZdmS9fqvHRCHbjUlO/+HNfgIwMWZ4FmsjD3wnPxONOm9YdVn/PRD7SsPRPbOjwBiE4EBGaHDsLjYAGDSGi7NJnSkA-----END"
                                            + " CERTIFICATE-----")
                                    .withLedgerRoleName(LedgerRoleName.READER))))
            .create();
    }

    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
