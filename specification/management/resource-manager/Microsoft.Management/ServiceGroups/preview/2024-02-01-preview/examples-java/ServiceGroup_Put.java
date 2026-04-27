
import com.azure.resourcemanager.servicegroups.fluent.models.ServiceGroupInner;
import com.azure.resourcemanager.servicegroups.models.ParentServiceGroupProperties;
import com.azure.resourcemanager.servicegroups.models.ServiceGroupProperties;

/**
 * Samples for ResourceProvider CreateOrUpdateServiceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-02-01-preview/ServiceGroup_Put.json
     */
    /**
     * Sample code: PutServiceGroup.
     * 
     * @param manager Entry point to ServiceGroupsManager.
     */
    public static void putServiceGroup(com.azure.resourcemanager.servicegroups.ServiceGroupsManager manager) {
        manager.resourceProviders().createOrUpdateServiceGroup("ServiceGroup1",
            new ServiceGroupInner().withProperties(new ServiceGroupProperties().withDisplayName("ServiceGroup 1 Name")
                .withParent(new ParentServiceGroupProperties()
                    .withResourceId("/providers/Microsoft.Management/serviceGroups/RootGroup"))),
            com.azure.core.util.Context.NONE);
    }
}
