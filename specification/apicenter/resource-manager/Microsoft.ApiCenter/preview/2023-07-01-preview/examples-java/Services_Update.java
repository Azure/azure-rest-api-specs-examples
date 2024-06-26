import com.azure.resourcemanager.apicenter.models.Service;

/** Samples for Services Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apicenter/resource-manager/Microsoft.ApiCenter/preview/2023-07-01-preview/examples/Services_Update.json
     */
    /**
     * Sample code: Services_Update.
     *
     * @param manager Entry point to ApiCenterManager.
     */
    public static void servicesUpdate(com.azure.resourcemanager.apicenter.ApiCenterManager manager) {
        Service resource =
            manager
                .services()
                .getByResourceGroupWithResponse("contoso-resources", "contoso", com.azure.core.util.Context.NONE)
                .getValue();
        resource.update().apply();
    }
}
