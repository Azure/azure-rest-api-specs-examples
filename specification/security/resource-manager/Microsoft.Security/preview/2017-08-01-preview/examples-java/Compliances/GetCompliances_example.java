import com.azure.core.util.Context;

/** Samples for Compliances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/Compliances/GetCompliances_example.json
     */
    /**
     * Sample code: Get security compliance data over time.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getSecurityComplianceDataOverTime(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.compliances().list("subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23", Context.NONE);
    }
}
