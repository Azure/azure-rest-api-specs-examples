
import com.azure.resourcemanager.storageactions.models.ElseCondition;
import com.azure.resourcemanager.storageactions.models.IfCondition;
import com.azure.resourcemanager.storageactions.models.OnFailure;
import com.azure.resourcemanager.storageactions.models.OnSuccess;
import com.azure.resourcemanager.storageactions.models.StorageTask;
import com.azure.resourcemanager.storageactions.models.StorageTaskAction;
import com.azure.resourcemanager.storageactions.models.StorageTaskOperation;
import com.azure.resourcemanager.storageactions.models.StorageTaskOperationName;
import com.azure.resourcemanager.storageactions.models.StorageTaskProperties;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for StorageTasks Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storageactions/resource-manager/Microsoft.StorageActions/stable/2023-01-01/examples/
     * storageTasksCrud/PatchStorageTask.json
     */
    /**
     * Sample code: PatchStorageTask.
     * 
     * @param manager Entry point to StorageActionsManager.
     */
    public static void patchStorageTask(com.azure.resourcemanager.storageactions.StorageActionsManager manager) {
        StorageTask resource = manager.storageTasks()
            .getByResourceGroupWithResponse("res4228", "mytask1", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new StorageTaskProperties().withEnabled(true).withDescription("My Storage task")
                .withAction(new StorageTaskAction()
                    .withIfProperty(new IfCondition().withCondition("[[equals(AccessTier, 'Cool')]]")
                        .withOperations(Arrays.asList(new StorageTaskOperation()
                            .withName(StorageTaskOperationName.SET_BLOB_TIER).withParameters(mapOf("tier", "Hot"))
                            .withOnSuccess(OnSuccess.CONTINUE).withOnFailure(OnFailure.BREAK))))
                    .withElseProperty(new ElseCondition().withOperations(
                        Arrays.asList(new StorageTaskOperation().withName(StorageTaskOperationName.DELETE_BLOB)
                            .withOnSuccess(OnSuccess.CONTINUE).withOnFailure(OnFailure.BREAK))))))
            .apply();
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
