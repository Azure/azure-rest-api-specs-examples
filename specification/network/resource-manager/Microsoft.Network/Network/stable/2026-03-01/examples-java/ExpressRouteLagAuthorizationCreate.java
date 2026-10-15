
import com.azure.resourcemanager.network.fluent.models.ExpressRouteLagAuthorizationInner;
import com.azure.resourcemanager.network.models.ExpressRouteLagAuthorizationPropertiesFormat;

/**
 * Samples for ExpressRouteLagAuthorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/ExpressRouteLagAuthorizationCreate.json
     */
    /**
     * Sample code: Create ExpressRouteLag Authorization.
     * 
     * @param manager Entry point to NetworkManager.
     */
    public static void createExpressRouteLagAuthorization(com.azure.resourcemanager.network.NetworkManager manager) {
        manager.serviceClient().getExpressRouteLagAuthorizations().createOrUpdate("rg1", "expressRouteLagName",
            "authorizationName",
            new ExpressRouteLagAuthorizationInner().withProperties(new ExpressRouteLagAuthorizationPropertiesFormat()),
            com.azure.core.util.Context.NONE);
    }
}
