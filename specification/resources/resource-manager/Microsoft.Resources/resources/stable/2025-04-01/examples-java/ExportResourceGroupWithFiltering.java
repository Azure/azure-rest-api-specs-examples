
import com.azure.resourcemanager.resources.models.ExportTemplateRequest;
import java.util.Arrays;

/**
 * Samples for ResourceGroups ExportTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-04-01/ExportResourceGroupWithFiltering.json
     */
    /**
     * Sample code: Export a resource group with filtering.
     * 
     * @param manager Entry point to ResourceManager.
     */
    public static void exportAResourceGroupWithFiltering(com.azure.resourcemanager.resources.ResourceManager manager) {
        manager.serviceClient().getResourceGroups().exportTemplate("my-resource-group",
            new ExportTemplateRequest().withResources(Arrays.asList(
                "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/my-resource-group/providers/My.RP/myResourceType/myFirstResource"))
                .withOptions("SkipResourceNameParameterization"),
            com.azure.core.util.Context.NONE);
    }
}
