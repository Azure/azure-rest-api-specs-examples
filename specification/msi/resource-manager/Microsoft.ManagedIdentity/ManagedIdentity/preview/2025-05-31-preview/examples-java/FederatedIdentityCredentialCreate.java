
import com.azure.resourcemanager.msi.fluent.models.FederatedIdentityCredentialInner;
import java.util.Arrays;

/**
 * Samples for FederatedIdentityCredentials CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FederatedIdentityCredentialCreate.json
     */
    /**
     * Sample code: FederatedIdentityCredentialCreate.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void federatedIdentityCredentialCreate(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().createOrUpdateWithResponse("rgName", "resourceName",
            "ficResourceName",
            new FederatedIdentityCredentialInner().withIssuer("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID")
                .withSubject("system:serviceaccount:ns:svcaccount")
                .withAudiences(Arrays.asList("api://AzureADTokenExchange")),
            com.azure.core.util.Context.NONE);
    }
}
