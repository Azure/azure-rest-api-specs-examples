import com.azure.core.util.Context;

/** Samples for ServicePrincipals Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getServicePrincipal.json
     */
    /**
     * Sample code: Get service principal.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getServicePrincipal(com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.servicePrincipals().getWithResponse(Context.NONE);
    }
}
