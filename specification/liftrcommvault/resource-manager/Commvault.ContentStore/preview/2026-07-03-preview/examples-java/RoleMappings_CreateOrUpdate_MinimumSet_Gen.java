
import com.azure.resourcemanager.commvaultcontentstore.fluent.models.RoleMappingInner;
import com.azure.resourcemanager.commvaultcontentstore.models.EntityInfo;
import com.azure.resourcemanager.commvaultcontentstore.models.EntityType;
import com.azure.resourcemanager.commvaultcontentstore.models.RoleAssignment;
import com.azure.resourcemanager.commvaultcontentstore.models.RoleMappingProperties;
import com.azure.resourcemanager.commvaultcontentstore.models.RoleName;
import java.util.Arrays;

/**
 * Samples for RoleMappings CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-03-preview/RoleMappings_CreateOrUpdate_MinimumSet_Gen.json
     */
    /**
     * Sample code: RoleMappings_CreateOrUpdate_MinimumSet - Single BackupAdmin role only.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void roleMappingsCreateOrUpdateMinimumSetSingleBackupAdminRoleOnly(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.roleMappings().createOrUpdateWithResponse("rgcommvault", "myCloudAccount",
            new RoleMappingInner().withProperties(new RoleMappingProperties()
                .withRoles(Arrays.asList(new RoleAssignment().withRoleName(RoleName.BACKUP_ADMIN)
                    .withEntities(Arrays.asList(new EntityInfo().withId("a1b2c3d4-e5f6-7890-abcd-ef1234567890")
                        .withDisplayName("Tenant Admins").withEntityType(EntityType.GROUP)))))),
            com.azure.core.util.Context.NONE);
    }
}
