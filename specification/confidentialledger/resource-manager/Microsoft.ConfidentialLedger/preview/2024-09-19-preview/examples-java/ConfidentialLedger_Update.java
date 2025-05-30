
import com.azure.resourcemanager.confidentialledger.models.AadBasedSecurityPrincipal;
import com.azure.resourcemanager.confidentialledger.models.CertBasedSecurityPrincipal;
import com.azure.resourcemanager.confidentialledger.models.ConfidentialLedger;
import com.azure.resourcemanager.confidentialledger.models.LedgerProperties;
import com.azure.resourcemanager.confidentialledger.models.LedgerRoleName;
import com.azure.resourcemanager.confidentialledger.models.LedgerType;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Ledger Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/confidentialledger/resource-manager/Microsoft.ConfidentialLedger/preview/2024-09-19-preview/
     * examples/ConfidentialLedger_Update.json
     */
    /**
     * Sample code: ConfidentialLedgerUpdate.
     * 
     * @param manager Entry point to ConfidentialLedgerManager.
     */
    public static void
        confidentialLedgerUpdate(com.azure.resourcemanager.confidentialledger.ConfidentialLedgerManager manager) {
        ConfidentialLedger resource = manager.ledgers().getByResourceGroupWithResponse("DummyResourceGroupName",
            "DummyLedgerName", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withTags(mapOf("additionProps2", "additional property value", "additionalProps1", "additional properties"))
            .withProperties(new LedgerProperties().withLedgerType(LedgerType.PUBLIC)
                .withAadBasedSecurityPrincipals(Arrays
                    .asList(new AadBasedSecurityPrincipal().withPrincipalId("34621747-6fc8-4771-a2eb-72f31c461f2e")
                        .withTenantId("bce123b9-2b7b-4975-8360-5ca0b9b1cd08")
                        .withLedgerRoleName(LedgerRoleName.ADMINISTRATOR)))
                .withCertBasedSecurityPrincipals(Arrays.asList(new CertBasedSecurityPrincipal().withCert(
                    "-----BEGIN CERTIFICATE-----\nMIIDUjCCAjqgAwIBAgIQJ2IrDBawSkiAbkBYmiAopDANBgkqhkiG9w0BAQsFADAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwHhcNMjAwOTIzMjIxODQ2WhcNMjEwOTIzMjIyODQ2WjAmMSQwIgYDVQQDExtTeW50aGV0aWNzIExlZGdlciBVc2VyIENlcnQwggEiMA0GCSqGSIb3DQEBAQUAA4IBDwAwggEKAoIBAQCX2s/Eu4q/eQ63N+Ugeg5oAciZua/YCJr41c/696szvSY7Zg1SNJlW88/nbz70+QpO55OmqlEE3QCU+T0Vl/h0Gf//n1PYcoBbTGUnYEmV+fTTHict6rFiEwrGJ62tvcpYgwapInSLyEeUzjki0zhOLJ1OfRnYd1eGnFVMpE5aVjiS8Q5dmTEUyd51EIprGE8RYAW9aeWSwTH7gjHUsRlJnHKcdhaK/v5QKJnNu5bzPFUcpC0ZBcizoMPAtroLAD4B68Jl0z3op18MgZe6lRrVoWuxfqnk5GojuB/Vu8ohAZKoFhQ6NB6r+LL2AUs+Zr7Bt26IkEdR178n9JMEA4gHAgMBAAGjfDB6MA4GA1UdDwEB/wQEAwIFoDAJBgNVHRMEAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAfBgNVHSMEGDAWgBS/a7PU9iOfOKEyZCp11Oen5VSuuDAdBgNVHQ4EFgQUv2uz1PYjnzihMmQqddTnp+VUrrgwDQYJKoZIhvcNAQELBQADggEBAF5q2fDwnse8egXhfaJCqqM969E9gSacqFmASpoDJPRPEX7gqoO7v1ww7nqRtRDoRiBvo/yNk7jlSAkRN3nRRnZLZZ3MYQdmCr4FGyIqRg4Y94+nja+Du9pDD761rxRktMVPSOaAVM/E5DQvscDlPvlPYe9mkcrLCE4DXYpiMmLT8Tm55LJJq5m07dVDgzAIR1L/hmEcbK0pnLgzciMtMLxGO2udnyyW/UW9WxnjvrrD2JluTHH9mVbb+XQP1oFtlRBfH7aui1ZgWfKvxrdP4zdK9QoWSUvRux3TLsGmHRBjBMtqYDY3y5mB+aNjLelvWpeVb0m2aOSVXynrLwNCAVA=\n-----END CERTIFICATE-----")
                    .withLedgerRoleName(LedgerRoleName.READER))))
            .apply();
    }

    // Use "Map.of" if available
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
