
import com.azure.resourcemanager.msi.fluent.models.FederatedIdentityCredentialInner;
import com.azure.resourcemanager.msi.models.ClaimsMatchingExpression;
import java.util.Arrays;

/**
 * Samples for FederatedIdentityCredentials CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/FlexibleFederatedIdentityCredentialCreate.json
     */
    /**
     * Sample code: FlexibleFederatedIdentityCredentialCreate.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void flexibleFederatedIdentityCredentialCreate(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getFederatedIdentityCredentials().createOrUpdateWithResponse("rgName", "resourceName",
            "ficResourceName",
            new FederatedIdentityCredentialInner().withIssuer("https://oidc.prod-aks.azure.com/TenantGUID/IssuerGUID")
                .withAudiences(Arrays.asList("api://AzureADTokenExchange"))
                .withClaimsMatchingExpression(new ClaimsMatchingExpression()
                    .withValue("claims['sub'] matches system:serviceaccount:ns:*").withLanguageVersion(1)),
            com.azure.core.util.Context.NONE);
    }
}
