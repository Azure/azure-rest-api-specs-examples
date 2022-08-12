import com.azure.core.util.Context;

/** Samples for ServicePrincipals List. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listServicePrincipalBySubscription.json
     */
    /**
     * Sample code: List service principal by subscription.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listServicePrincipalBySubscription(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.servicePrincipals().list(Context.NONE);
    }
}
