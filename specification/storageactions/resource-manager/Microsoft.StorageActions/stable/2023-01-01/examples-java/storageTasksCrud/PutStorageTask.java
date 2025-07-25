
import com.azure.resourcemanager.storageactions.models.ElseCondition;
import com.azure.resourcemanager.storageactions.models.IfCondition;
import com.azure.resourcemanager.storageactions.models.ManagedServiceIdentity;
import com.azure.resourcemanager.storageactions.models.ManagedServiceIdentityType;
import com.azure.resourcemanager.storageactions.models.OnFailure;
import com.azure.resourcemanager.storageactions.models.OnSuccess;
import com.azure.resourcemanager.storageactions.models.StorageTaskAction;
import com.azure.resourcemanager.storageactions.models.StorageTaskOperation;
import com.azure.resourcemanager.storageactions.models.StorageTaskOperationName;
import com.azure.resourcemanager.storageactions.models.StorageTaskProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageTasks Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-01-01/storageTasksCrud/PutStorageTask.json
     */
    /**
     * Sample code: PutStorageTask.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void putStorageTask(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        manager.storageTasks().define("mytask1").withRegion("westus").withExistingResourceGroup("res4228")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED))
            .withProperties(new StorageTaskProperties().withEnabled(true).withDescription("My Storage task")
                .withAction(new StorageTaskAction()
                    .withIfProperty(new IfCondition().withCondition("[[equals(AccessTier, 'Cool')]]")
                        .withOperations(Arrays.asList(new StorageTaskOperation()
                            .withName(StorageTaskOperationName.SET_BLOB_TIER).withParameters(mapOf("tier", "Hot"))
                            .withOnSuccess(OnSuccess.CONTINUE).withOnFailure(OnFailure.BREAK))))
                    .withElseProperty(new ElseCondition().withOperations(
                        Arrays.asList(new StorageTaskOperation().withName(StorageTaskOperationName.DELETE_BLOB)
                            .withOnSuccess(OnSuccess.CONTINUE).withOnFailure(OnFailure.BREAK))))))
            .create();
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
