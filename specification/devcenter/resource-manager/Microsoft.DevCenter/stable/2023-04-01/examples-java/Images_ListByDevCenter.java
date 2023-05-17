/** Samples for Images ListByDevCenter. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/stable/2023-04-01/examples/Images_ListByDevCenter.json
     */
    /**
     * Sample code: Images_ListByDevCenter.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void imagesListByDevCenter(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager.images().listByDevCenter("rg1", "Contoso", null, com.azure.core.util.Context.NONE);
    }
}
