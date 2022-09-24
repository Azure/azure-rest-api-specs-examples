import com.azure.core.util.Context;

/** Samples for Applications List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-07-01-preview/examples/Applications/ListBySubscriptionApplications_example.json
     */
    /**
     * Sample code: List applications security by subscription level scope.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void listApplicationsSecurityBySubscriptionLevelScope(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.applications().list(Context.NONE);
    }
}
