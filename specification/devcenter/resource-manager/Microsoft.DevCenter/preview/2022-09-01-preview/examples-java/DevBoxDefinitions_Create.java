import com.azure.resourcemanager.devcenter.models.ImageReference;
import com.azure.resourcemanager.devcenter.models.Sku;

/** Samples for DevBoxDefinitions CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/DevBoxDefinitions_Create.json
     */
    /**
     * Sample code: DevBoxDefinitions_Create.
     *
     * @param manager Entry point to DevCenterManager.
     */
    public static void devBoxDefinitionsCreate(com.azure.resourcemanager.devcenter.DevCenterManager manager) {
        manager
            .devBoxDefinitions()
            .define("WebDevBox")
            .withRegion("centralus")
            .withExistingDevcenter("rg1", "Contoso")
            .withImageReference(
                new ImageReference()
                    .withId(
                        "/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/Example/providers/Microsoft.DevCenter/devcenters/Contoso/galleries/contosogallery/images/exampleImage/version/1.0.0"))
            .withSku(new Sku().withName("Preview"))
            .withOsStorageType("SSD_1024")
            .create();
    }
}
