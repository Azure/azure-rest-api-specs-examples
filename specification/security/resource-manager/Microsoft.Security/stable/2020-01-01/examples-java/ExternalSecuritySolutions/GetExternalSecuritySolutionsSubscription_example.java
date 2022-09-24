import com.azure.core.util.Context;

/** Samples for ExternalSecuritySolutions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ExternalSecuritySolutions/GetExternalSecuritySolutionsSubscription_example.json
     */
    /**
     * Sample code: Get external security solutions on a subscription.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getExternalSecuritySolutionsOnASubscription(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.externalSecuritySolutions().list(Context.NONE);
    }
}
