
/**
 * Samples for TrustedAccessRoles List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/TrustedAccessRoles_List.json
     */
    /**
     * Sample code: List trusted access roles.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listTrustedAccessRoles(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getTrustedAccessRoles().list("westus2", com.azure.core.util.Context.NONE);
    }
}
