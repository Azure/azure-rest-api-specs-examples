
import com.azure.resourcemanager.compute.fluent.models.GalleryInVMAccessControlProfileVersionInner;
import com.azure.resourcemanager.compute.models.AccessControlRules;
import com.azure.resourcemanager.compute.models.AccessControlRulesIdentity;
import com.azure.resourcemanager.compute.models.AccessControlRulesMode;
import com.azure.resourcemanager.compute.models.AccessControlRulesPrivilege;
import com.azure.resourcemanager.compute.models.AccessControlRulesRole;
import com.azure.resourcemanager.compute.models.AccessControlRulesRoleAssignment;
import com.azure.resourcemanager.compute.models.EndpointAccess;
import com.azure.resourcemanager.compute.models.TargetRegion;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for GalleryInVMAccessControlProfileVersions CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfileVersion_Create.json
     */
    /**
     * Sample code: Create or update a Gallery InVMAccessControlProfile Version.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        createOrUpdateAGalleryInVMAccessControlProfileVersion(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfileVersions().createOrUpdate(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName", "1.0.0",
            new GalleryInVMAccessControlProfileVersionInner().withLocation("West US")
                .withMode(AccessControlRulesMode.AUDIT).withDefaultAccess(EndpointAccess.ALLOW)
                .withRules(new AccessControlRules()
                    .withPrivileges(Arrays.asList(new AccessControlRulesPrivilege().withName("GoalState")
                        .withPath("/machine").withQueryParameters(mapOf("comp", "goalstate"))))
                    .withRoles(Arrays.asList(new AccessControlRulesRole().withName("Provisioning")
                        .withPrivileges(Arrays.asList("GoalState"))))
                    .withIdentities(Arrays.asList(new AccessControlRulesIdentity().withName("WinPA")
                        .withUsername("SYSTEM").withGroupName("Administrators")
                        .withExePath("C:\\Windows\\System32\\cscript.exe").withProcessName("cscript")))
                    .withRoleAssignments(Arrays.asList(new AccessControlRulesRoleAssignment().withRole("Provisioning")
                        .withIdentities(Arrays.asList("WinPA")))))
                .withTargetLocations(Arrays.asList(new TargetRegion().withName("West US"),
                    new TargetRegion().withName("South Central US")))
                .withExcludeFromLatest(false),
            com.azure.core.util.Context.NONE);
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
