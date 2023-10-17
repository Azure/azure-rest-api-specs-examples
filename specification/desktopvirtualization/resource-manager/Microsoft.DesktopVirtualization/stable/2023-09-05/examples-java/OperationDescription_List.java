/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/OperationDescription_List.json
     */
    /**
     * Sample code: OperationDescription_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void operationDescriptionList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
